<?php

$date = [
    '2019-12-31 12:22:33',
    '2019-11-10 12:22:22',
    '2018-12-01 12:22:33',
	'2021-11-05 11:11:11',
    '2019-05-11 11:11:11'
];

$com = time();

$cnt = count($date);

for($i=0 ; $i < $cnt; $i++){
        if(isset($date[$i+1])) {           
            $front = abs($com - strtotime($date[$i]));
            $behind = abs($com - strtotime($date[$i+1]));          
            if($front > $behind) {
                list($date[$i], $date[$i+1]) = [$date[$i+1], $date[$i]];
            }
        }
    }

var_dump($date);


// 单一原则、开放封闭原则、李氏代替原则、接口分离原则、依赖倒置原则
// 三范式：属性原子性：不能在拆分
// 三范式：数据唯一性：数据不能在多个表中存储
// 三范式：字段冗余性：任何字段都不能派生其它字段
