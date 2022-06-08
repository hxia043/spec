package hash

/*
#include <stdio.h>
#include <stddef.h>
#include <crypt.h>
#include <string.h>
#include <stdbool.h>
#include <stdlib.h>

#cgo LDFLAGS: -lcrypt

char *makeHashPassword(const char* passwd,int rounds){
    struct crypt_data data;
    memset(&data, 0, sizeof(data));
    if (crypt_gensalt_rn("$6$", rounds, NULL, 0, data.setting, sizeof(data.setting)) == NULL)
    {
        return NULL;
    }
    if (strlen(passwd) >= sizeof(data.input))
    {
        return NULL;
    }
    snprintf(data.input, sizeof(data.input), "%s", passwd);
    char* hash = crypt_rn(data.input, data.setting, &data, sizeof(data));
    if (hash == NULL)
    {
       return NULL;
    }
    char* ret = strdup(hash);
    memset(&data, 0, sizeof(data));
    return ret;
}

bool isPasswdMatches(const char *passwd, const char* hash)
{
    struct crypt_data data;
    memset(&data, 0, sizeof(data));
    if (strlen(passwd) >= sizeof(data.input))
        return false;
    if (strlen(hash) >= sizeof(data.setting))
        return false;
    snprintf(data.input, sizeof(data.input), "%s", passwd);
    snprintf(data.setting, sizeof(data.setting), "%s", hash);
    char* newhash = crypt_rn(data.input, data.setting, &data, sizeof(data));
    if (newhash == NULL)
        return false;
    int ok = strcmp(hash, newhash) == 0;
    memset(&data, 0, sizeof(data));
    return ok;
}

*/
import "C"
import "unsafe"

func HashPassword(password string) (string, error) {
	hash, err := C.makeHashPassword(C.CString(password), C.int(1000))
	defer C.free(unsafe.Pointer(hash))
	return C.GoString(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	return bool(C.isPasswdMatches(C.CString(password), C.CString(hash)))
}
