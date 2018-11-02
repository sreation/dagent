package common

import (
    "io/ioutil"
    "encoding/json"
)

func ShowFile(myPath string) (string, bool){
    files, err := ioutil.ReadDir(myPath)
    if err != nil {
        Error.Println("operation:showfile "+myPath+", fail.")
        return "", false
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
        Error.Println("operation:showfile jsonencode map fail.")
        return "", false
    }
    return string(data), true
}

func ReadFile(filePth string) (string, bool) {
    result, err := ioutil.ReadFile(filePth)
    if err != nil {
        Error.Println("operation:readfile "+filePth+" is fail.")
        return "", false
    }
    return string(result), true
}
