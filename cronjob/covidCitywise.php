<?php
$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => "https://api.covid19india.org/v2/state_district_wise.json",
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

$cityData = json_decode($response,true);
echo "data--->".sizeof($cityData);
/*
key1 :: district ---> Value :: South Andaman
key1 :: active ---> Value :: 5
key1 :: confirmed ---> Value :: 15
key1 :: deceased ---> Value :: 0
key1 :: recovered ---> Value :: 10
 */

for ($i=0;$i<=sizeof($cityData);$i++){
  foreach($cityData[$i] as $k => $v){
    if ($k == "districtData"){
      for ($j=0;$j<=sizeof($v);$j++){
        foreach($v[$j] as $ink => $inv){
          `redis-cli del tempDistrictKey`;
          if ($ink == "district"){
            $redisKey = strtolower($inv);
            $redisKey = str_replace(' ',"-",$redisKey); 
            echo "\n----------".$redisKey."\n";
            echo "redis-cli set tempDistrictKey $redisKey";
          }
          if (($ink == "active") || ($ink == "confirmed") || ($ink == "deceased") || ($ink == "recovered")){
            $myRedisVal = `redis-cli get tempDistrictKey`;  
            echo "redis-cli hset $myRedisVal $ink $inv"; 
          }
        }
        exit;
      }
    }
  }
  exit;   
}
exit;
var_dump($cityData[1]);
echo "Done";

?>
