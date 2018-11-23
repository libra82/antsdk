package utils

import (
  "net/url"
  "sort"
)


func BuildQuery(params map[string]string) string {
  if params == nil {
    return ""
  }

  query := NewStringBuilder()
  hasParam := false

  for k, v := range params {
    if v != "" {
      if hasParam {
        query.Append("&")
      } else {
        hasParam = true
      }

      query.Append(k).Append("=").Append(url.QueryEscape(v))
    }
  }
  return query.ToString()
}

func BuildAndSortQuery(params map[string]string) string {
  if params == nil {
    return ""
  }

  // 对 key 进行升序排序
  sortedKeys := make([]string, 0)
  for k, _ := range params {
    sortedKeys = append(sortedKeys, k)
  }
  sort.Strings(sortedKeys)

  // 对 key=value 的键值对用 & 连接起来，略过空值
  query := NewStringBuilder()
  for i, k := range sortedKeys {
    if params[k] != "" {
      query.Append(k)
      query.Append("=")
      query.Append(url.QueryEscape(params[k]))
      if i != (len(sortedKeys) - 1) {
        query.Append("&")
      }
    }
  }

  return query.ToString()
}

func MergeMap(m1, m2, m3 map[string]string) (map[string]string, string) {
  var retMap = make(map[string]string)
  var sign string
  for key, value := range m1 {
    if key == "sign" {
      sign = value
      continue
    }
    retMap[key] = value
  }
  for key, value := range m2 {
    if key == "sign" {
      sign = value
      continue
    }
    retMap[key] = value
  }
  for key, value := range m3 {
    if key == "sign" {
      sign = value
      continue
    }
    retMap[key] = value
  }

  return retMap, sign
}