package password

import "testing"

func TestEncode(t *testing.T) {
	code := Encode("mypasswd", 64, "default")
	t.Log(code)
}

//pbkdf2_sha256$216000$wRhTGOBVdwGZ97yZGGzYIO48L6c6J8gZ3woXn8NZyCJQbWgpLUWJnUJPMtddtGNE$WCLgOmYqqQtTrmfWR3c0q3KHdnamRCUirmyoGiaMlvE=
//pbkdf2_sha256$216000$W4wvu39TmOoy6JznbiPUbNiw2GPLxg58URr1woecBPejH1WFQs1M6xk5oeZ5d7LH$LsDSWUC6Q30dfZ4+AYhn9B1ndHf022IOnE8Ae8tZuCU=
func TestValidate(t *testing.T) {
	ok := Validate("mypasswd", "pbkdf2_sha256$216000$wRhTGOBVdwGZ97yZGGzYIO48L6c6J8gZ3woXn8NZyCJQbWgpLUWJnUJPMtddtGNE$WCLgOmYqqQtTrmfWR3c0q3KHdnamRCUirmyoGiaMlvE=")
	t.Log(ok)
}
