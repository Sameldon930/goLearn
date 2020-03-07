<?php

$arr = [
    [
        'a'=>1,
        'b'=>2
    ],
    [
        'c'=>3,
        'd'=>4
    ],
    [
        'e'=>5,
        'f'=>6
    ]
];
foreach ($arr as $v1){
    foreach($v1 as $v2){
        print_r($v2."\n");
    }
}