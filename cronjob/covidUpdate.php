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
$covidData = json_decode($response, true);

//setting up Global status
foreach($covidData["Global"] as $k => $v){
  `redis-cli hset globalStatus $k $v`;
  echo "Key :: $k ===> Value :: $v \n";
}
curl_close($curl);
