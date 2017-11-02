package ini

import (
	"fmt"
	"os"
	"testing"
)

func TestStringWithoutEnv(t *testing.T) {
	s := "this test without env"
	if s != replace_env(s) {
		t.Fail()
	}
}

func TestStringWithEnv(t *testing.T) {
	s := "this is my ${HOME} env"
	new_value := fmt.Sprintf("this is my %s env", os.Getenv("HOME"))
	if new_value != replace_env(s) {
		t.Fail()
	}
}

func TestStringWithEnvDefaultValue(t *testing.T) {
	s := "there is no such var ${MY_NOT_EXIST_ENV:-test},haha"
	if "there is no such var test,haha" != replace_env(s) {
		t.Fail()
	}
}

func TestStringWithEscapedBeforeEnv(t *testing.T) {
	s := "this is my \\${HOME} env"
	if s != replace_env(s) {
		t.Fail()
	}
}
