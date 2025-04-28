// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    // 存储键值对，现在键为 uint256 类型，值为字符串类型
    mapping(uint256 => string) private keyValue;

    // 设置字符串类型的值
    function set(uint256 key, string memory value) public {
        keyValue[key] = value;
    }

    // 获取字符串类型的值
    function get(uint256 key) public view returns (string memory) {
        return keyValue[key];
    }
}