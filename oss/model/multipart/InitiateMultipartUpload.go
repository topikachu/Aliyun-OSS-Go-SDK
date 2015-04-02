/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"Aliyun-OSS-Go-SDK/oss/types"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// 	Start the multipartUpload and  get the UploadId.
/*
 *	Example:
 *	initObjectPath, imur, err := c.InitiateMultipartUpload("bucketName/test.txt")
 */
func (c *Client) InitiateMultipartUpload(objectPath string) (initObjectPath string, imur types.InitiateMultipartUploadResult, err error) {
	cc := ConvertClient(c)

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	initObjectPath = objectPath
	resp, err := cc.DoRequest("POST", objectPath+"?uploads", objectPath+"?uploads", nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &imur)
	fmt.Println("The multipart upload has been intiated and you have got the UploadId.")
	return
}