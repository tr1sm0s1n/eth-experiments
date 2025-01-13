// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract Datastore {
    string public data;
    struct EventInfo {
        uint start;
        uint end;
        bool first;
    }
    string[][] stringArray;

    mapping(string => string[][]) stringMap;
    mapping(string => uint) datacount;
    mapping(string => EventInfo) public eventCount;

    event DataStored(string indexed exam_no, string[] data);

    function setData(string memory _data) public {
        data = _data;
    }

    function storeData(string[][] memory newStrings) public {
        stringArray = newStrings;
    }

    function storeDataWithEvent(string[][] memory newStrings) public {
        for (uint i = 0; i < newStrings.length; i++) {
            stringArray.push(newStrings[i]);
            emit DataStored(newStrings[i][3], newStrings[i]);
        }
    }

    function storeDataMap(string[][] memory newStrings) public {
        for (uint i = 0; i < newStrings.length; i++) {
            stringMap[newStrings[i][3]].push(newStrings[i]);
            emit DataStored(newStrings[i][3], newStrings[i]);
            eventCount[newStrings[i][3]].end = block.number;

            if (!eventCount[newStrings[i][3]].first) {
                eventCount[newStrings[i][3]].start = block.number;
                eventCount[newStrings[i][3]].first = true;
            }
        }
    }

    function getAllData() public view returns (string[][] memory) {
        return stringArray;
    }

    function getDataArray(uint index) public view returns (string[] memory) {
        return stringArray[index];
    }

    function getDataMap(
        string memory _exam_no,
        uint _start,
        uint _number
    ) public view returns (string[][] memory) {
        string[][] storage dataMap = stringMap[_exam_no];
        require(_start + _number <= dataMap.length, "Out of bounds");

        string[][] memory batch = new string[][](_number);
        for (uint i = 0; i < _number; i++) {
            batch[i] = dataMap[_start + i];
        }
        return batch;
    }
}
