package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"sigs.k8s.io/yaml"
)

type Meta struct {
	Source  string
	Version string
}

func Load(input string) (*openapi3.T, *Meta, error) {
	data, source, err := readInput(input)
	if err != nil {
		return nil, nil, err
	}

	jsonData, err := toJSON(data)
	if err != nil {
		return nil, nil, err
	}

	version, err := detectVersion(jsonData)
	if err != nil {
		return nil, nil, err
	}

	if version == "openapi3" {
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = true
		doc, err := loader.LoadFromData(jsonData)
		if err != nil {
			return nil, nil, fmt.Errorf("load openapi3 failed: %w", err)
		}
		return doc, &Meta{Source: source, Version: "OpenAPI 3"}, nil
	}

	var doc2 openapi2.T
	if err := json.Unmarshal(jsonData, &doc2); err != nil {
		return nil, nil, fmt.Errorf("load swagger2 failed: %w", err)
	}
	doc3, err := openapi2conv.ToV3(&doc2)
	if err != nil {
		return nil, nil, fmt.Errorf("convert swagger2 to openapi3 failed: %w", err)
	}

	return doc3, &Meta{Source: source, Version: "Swagger 2.0"}, nil
}

func readInput(input string) ([]byte, string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return nil, "", errors.New("input is empty")
	}

	if strings.HasPrefix(trimmed, "http://") || strings.HasPrefix(trimmed, "https://") {
		client := &http.Client{Timeout: 20 * time.Second}
		resp, err := client.Get(trimmed)
		if err != nil {
			return nil, "", fmt.Errorf("fetch url failed: %w", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, "", fmt.Errorf("fetch url failed: status %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, "", fmt.Errorf("read url response failed: %w", err)
		}
		return body, trimmed, nil
	}

	data, err := os.ReadFile(trimmed)
	if err != nil {
		return nil, "", fmt.Errorf("read file failed: %w", err)
	}
	return data, trimmed, nil
}

func toJSON(data []byte) ([]byte, error) {
	var probe map[string]any
	if err := json.Unmarshal(data, &probe); err == nil {
		return data, nil
	}

	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, fmt.Errorf("convert yaml to json failed: %w", err)
	}

	return jsonData, nil
}

func detectVersion(jsonData []byte) (string, error) {
	var probe map[string]any
	if err := json.Unmarshal(jsonData, &probe); err != nil {
		return "", fmt.Errorf("parse json failed: %w", err)
	}

	if _, ok := probe["openapi"]; ok {
		return "openapi3", nil
	}
	if _, ok := probe["swagger"]; ok {
		return "swagger2", nil
	}

	return "", errors.New("cannot detect spec version: missing openapi or swagger")
}
