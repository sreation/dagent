package common

import (
    "io/ioutil"
    "encoding/json"
    "os"
)

func ShowFile(myPath string) (bool, string, string){
    files, err := ioutil.ReadDir(myPath)
    if err != nil {
        Error.Println("operation:showfile " + err.Error())
        return false, err.Error(), ""
    }

    fileMap := []map[string]interface{}{}
    for _, f := range files {
        x := map[string]interface{}{
            "name": f.Name(),
            "mode": f.Mode().String(),
            "modtime": f.ModTime().String(),
        }
        fileMap = append(fileMap, x)
    }

    data, err := json.Marshal(fileMap)
    if err != nil {
        Error.Println("operation:showfile-json " + err.Error())
        return false, err.Error(), ""
    }
    return true, "", string(data)
}

//func FileExist(path string) (bool, bool) {
//    _, err := os.Stat(path)
//    if err == nil {
//        return true, true
//    }
//    if os.IsNotExist(err) {
//        return false, true
//    }
//    Error.Println("operation:fileexist "+path+", fail.")
//    return false, false
//}

func ReadFile(filePath string) (bool, string, string) {
    result, err := ioutil.ReadFile(filePath)
    if err != nil {
        Error.Println("operation:readfile " + err.Error())
        return false, err.Error(), ""
    }
    return true, "", string(result)
}

func WriteFile(filePath string, fileContent string) (bool, string) {
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
    //file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE , os.ModePerm)
    defer file.Close()
    if err != nil {
        Error.Println("operation:WriteFile " + err.Error())
        return false, err.Error()
    } else {
        bo, s := JsonDecode(fileContent)
        if bo == false {
            return false, s
        } else {
            _, err = file.Write([]byte(s))
            if err != nil {
                return false, err.Error()
            }
        }
    }
    return true, ""
}
