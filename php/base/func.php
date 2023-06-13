<?php
// 保存一些实用的php函数

/**
 * 转换为英语复数
 * @param string $word 需要转换的单词
 * @return string
 */
function pluralize($word)
{
    $rules = array(
        '/move$/i' => 'moves',
        '/foot$/i' => 'feet',
        '/child$/i' => 'children',
        '/human$/i' => 'humans',
        '/man$/i' => 'men',
        '/tooth$/i' => 'teeth',
        '/person$/i' => 'people',
        '/([m|l])ouse$/i' => '\1ice',
        '/(x|ch|ss|sh|us|as|is|os)$/i' => '\1es',
        '/([^aeiouy]|qu)y$/i' => '\1ies',
        '/(?:([^f])fe|([lr])f)$/i' => '\1\2ves',
        '/(shea|lea|loa|thie)f$/i' => '\1ves',
        '/([ti])um$/i' => '\1a',
        '/(tomat|potat|ech|her|vet)o$/i' => '\1oes',
        '/(bu)s$/i' => '\1ses',
        '/(ax|test)is$/i' => '\1es',
        '/s$/' => 's',
    );
    foreach ($rules as $rule => $replacement) {
        if (preg_match($rule, $word))
            return preg_replace($rule, $replacement, $word);
    }
    return $word . 's';
}

/**
 * 转array
 * @param $data
 * @return array
 */
function to_array($data)
{
    $data = json_encode($data, JSON_UNESCAPED_UNICODE);
    $data = json_decode($data, true);

    return $data;
}

/**
 * 检测并删除json里面的包含文件
 * @param json $data
 */
function unlink_json_files($data)
{
    // 转php数组
    $data = json_decode($data, true);
    foreach ($data as $v) {
        if (is_array($v)) {
            // 递归
            unlink_json_files(json_encode($v, JSON_UNESCAPED_UNICODE));
        } else if (is_string($v)) {
            if (_is_file($v)) {
                // 删除
                _unlink($v);
                continue;
            }
        }
    }
}

/**
 * 判断是否是移动端
 * @return bool
 */
function is_mobile()
{
    return strrpos($_SERVER['HTTP_USER_AGENT'], "Mobile") ? true : false;
}

/**
 * 判断苹果或者安卓设备
 * @return string
 */
function mobile_type()
{
    $Android = "Android";
    $iPhone = "iPhone";
    if (strrpos($_SERVER['HTTP_USER_AGENT'], "Android")) {
        return $Android;
    }
    if (strrpos($_SERVER['HTTP_USER_AGENT'], "iPhone")) {
        return $iPhone;
    }
}

/**
 * http请求
 * @param string $url 请求url
 * @param array|null $postData post请求数组
 * @return data
 */
function http_request($url, $postData = null)
{
    // 初始化curl
    $curl = curl_init();
    // 设置需要抓取的url
    curl_setopt($curl, CURLOPT_URL, $url);
    // 设置头文件的信息作为数据流输出
    // curl_setopt($curl, CURLOPT_HEADER, 1);

    // 设置证书
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
    // 设置获取的信息以文件流的形式返回，而不是直接输出
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);

    if ($postData) {
        // 设置post方式提交
        curl_setopt($curl, CURLOPT_POST, 1);
        // 设置post数据
        curl_setopt($curl, CURLOPT_POSTFIELDS, $postData);
    }
    // 执行请求
    $data = curl_exec($curl);
    curl_close($curl);

    return $data;
}

/**
 * 验证是否是正确的手机号码
 * @param $phone 手机号码
 * @return bool
 * @return bool
 */
function is_phone($phone)
{
    $pattern = "/^1[34578]\d{9}$/";
    preg_match($pattern, $phone, $msg);
    return $msg ? true : false;
}

/**
 * 用*代替指定的字符
 * @param string $str 字符串
 * @param int $start 开始位置
 * @param int $length 要替换的长度
 * @return string|string[]
 * @return string|string[]
 */
function hide_char($str, $start, $length)
{
    // 根据长度确定*的数量
    $hide = "";
    for ($i = 0; $i < $length; $i++) {
        $hide .= "*";
    }
    $str = substr_replace($str, $hide, --$start, $length);
    return $str;
}

/**
 * 对二维数组根据某一个键值进行排序
 * @param array $array 二维数组
 * @param string $keys 需要进行排序的键
 * @param string $sort 排序方式，默认asc,可选desc
 * @return array
 */
function array_sort($array, $key, $sort = 'asc')
{
    if (is_array($array)) {
        if (count($array) == count($array, 1)) {
            return [];
        }
    } else {
        return [];
    }
    $newArray = $tempArray = [];

    // 提出key的一维数组
    foreach ($array as $k => $v) {
        $tempArray[$k] = $v[$key];
    }
    // 对一维数组进行排序
    $sort == 'asc' ? asort($tempArray) : arsort($tempArray);

    // 重新组成新的数组
    foreach ($tempArray as $k => $v) {
        $newArray[$k] = $array[$k];
    }

    // 重置键的顺序并返回
    $array = array_merge($newArray);
    return $array;
}

/**
 * base64格式编码转换为图片并保存对应文件夹
 * @param string $base64_image_content base64内容
 * @param string $path 保存的文件夹
 * @return bool|false|string
 * @return bool|false|string
 */
function base64_image_content($base64_image_content, $path)
{
    // 匹配出图片的格式
    if (preg_match('/^(data:\s*image\/(\w+);base64,)/', $base64_image_content, $result)) {
        $type = $result[2];
        $newFile = "." . $path . "/";
        // 检测是否有该文件夹,没有则创建
        if (!is_dir($newFile)) {
            mkdir($newFile, 0777, true);
        }
        $path = $newFile . time() . ".{$type}";
        if (file_put_contents($path, base64_decode(str_replace($result[1], '', $base64_image_content)))) {
            return substr($path, 1);
        } else {
            return false;
        }
    } else {
        return false;
    }
}

/**
 * 读取一个文件夹下的所有文件
 * @param string $dir 文件夹路径
 * @return array
 * @return array
 */
function scan_dir($dir)
{
    // 定义一个数组
    $files = array();
    // 检测是否存在文件
    if (is_dir($dir)) {
        // 打开目录
        if ($handle = opendir($dir)) {
            // 返回当前文件的条目
            while (($file = readdir($handle)) !== false) {
                // 去除特殊目录
                if ($file != "." && $file != "..") {
                    // 判断子目录是否还存在子目录
                    if (is_dir($dir . "/" . $file)) {
                        // 递归调用本函数，再次获取目录
                        $files[$file] = scan_dir($dir . "/" . $file);
                    } else {
                        // 获取目录数组
                        $files[] = $dir . "/" . $file;
                    }
                }
            }
            // 关闭文件夹
            closedir($handle);
            // 返回文件夹数组
            return $files;
        }
    }
}

/**
 * 批量修改一个目录下的所有文件名
 * @param string $dir 目录
 * @param string $name 名称
 * @return bool
 */
function batch_file_name($dir, $name)
{
    // 读取文件下的名称
    $data = scan_dir($dir);
    if (is_array($data)) {
        foreach ($data as $v) {
            // 获取基础信息
            $info = pathinfo($v);
            $basename = $info['basename'];
            $extension = $info['extension'];
            $newFile = $name . '.' . $extension;

            // 重新命名
            $filePath = $dir . '/' . $basename;
            $newFilePath = $dir . '/' . $newFile;
            rename($filePath, $newFilePath);
        }
    } else {
        return false;
    }
}

/**
 * 高维数组降低为一维数组
 * @param $array 数组
 * @param bool $unique 是否去除重复元素，默认去除
 * @return array
 */
function flareout_array($array, $unique = true)
{
    // 拉平
    $arr = [];
    array_walk_recursive($array, function ($x) use (&$arr) {
        $arr[] = $x;
    });

    if ($unique) {
        // 去重
        $arr = array_values(array_unique($arr));
    }

    return $arr;
}

/**
 * xml字符串转数组
 */
function xml_array($xml)
{
    // 将字符串转化为变量
    $data = simplexml_load_string($xml, 'SimpleXMLElement', LIBXML_NOCDATA);
    // 转JSON
    $json = json_encode($data, JSON_UNESCAPED_UNICODE);
    // 返回数组
    return json_decode($json, true);
}

/**
 * 判断远程文件是否存在
 * @param string $url 文件url地址
 * @return bool
 */
function _is_file($url)
{
    if (is_file('.' . parse_url($url)['path'])) {
        return true;
    } else {
        return false;
    }
}

/**
 * 删除以远程文件当地址的文件
 * @param string $url 文件url地址
 */
function _unlink($url) {
    unlink('.' . parse_url($url)['path']);
}
