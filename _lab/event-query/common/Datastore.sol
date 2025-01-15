// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract DataStore {
    event Stored(string indexed exam_no, string[] data);

    struct EventInfo {
        uint start;
        uint end;
        bool first;
    }

    mapping(string => EventInfo) public EventCount;

    function storeData(string[][] memory newStrings) public {
        for (uint i = 0; i < newStrings.length; i++) {
            emit Stored(newStrings[i][3], newStrings[i]);
            EventCount[newStrings[i][3]].end = block.number;

            if (!EventCount[newStrings[i][3]].first) {
                EventCount[newStrings[i][3]].start = block.number;
                EventCount[newStrings[i][3]].first = true;
            }
        }
    }
}
