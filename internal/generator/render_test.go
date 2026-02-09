package generator

import (
	"strings"
	"testing"
)

func TestRenderAPIFile_ModelImportsWithoutTrailingComma(t *testing.T) {
	content := RenderAPIFile(nil, []string{"SendEmailCodeForm", "Captcha", "GetCaptchaContentParam"}, false)

	expectedImportBlock := "import type {\n  Captcha,\n  GetCaptchaContentParam,\n  SendEmailCodeForm\n} from './model';\n"
	if !strings.Contains(content, expectedImportBlock) {
		t.Fatalf("unexpected model import block:\n%s", content)
	}

	unexpectedImportBlock := "SendEmailCodeForm,\n} from './model';"
	if strings.Contains(content, unexpectedImportBlock) {
		t.Fatalf("model import contains trailing comma before closing brace:\n%s", content)
	}
}

func TestRenderAPIFile_SingleModelImportWithoutTrailingComma(t *testing.T) {
	content := RenderAPIFile(nil, []string{"LoginForm"}, false)

	expectedImportBlock := "import type {\n  LoginForm\n} from './model';\n"
	if !strings.Contains(content, expectedImportBlock) {
		t.Fatalf("unexpected single model import block:\n%s", content)
	}

	unexpectedImportBlock := "LoginForm,\n} from './model';"
	if strings.Contains(content, unexpectedImportBlock) {
		t.Fatalf("single model import contains trailing comma before closing brace:\n%s", content)
	}
}
