<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => "https://api.covid19api.com/summary",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "GET",
));

$response = curl_exec($curl);
curl_close($curl);

$covidData = json_decode($response, true);
$countries_data = $covidData["Countries"];

//setting up Global status
foreach($covidData["Global"] as $k => $v){
  `redis-cli hset globalStatus $k $v`;
  echo "Key :: $k ===> Value :: $v \n";
}

$maxLen = sizeof($countries_data);

/*
| ====  key for hash to be stored in Redis ==== |
              NewConfirmed 
              TotalConfirmed
              NewDeaths
              TotalDeaths
              NewRecovered 
              TotalRecovered
| ============================================= |
 */ 

for ($i = 0; $i <= $maxLen; $i++){
  //Getting Countries data
  foreach($countries_data[$i] as $k => $v){
    if($k == "Slug"){
      $myRedisKey = $v;
    }

    if ($k == "NewConfirmed" || $k == "TotalConfirmed" || $k =="NewDeaths" || $k == "TotalDeaths" || $k == "NewRecovered" || $k == "TotalRecovered"){
      `redis-cli hset $myRedisKey $k $v`;
    }
  }
}
