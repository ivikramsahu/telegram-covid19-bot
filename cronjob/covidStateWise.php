<?php
$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => "https://ameerthehacker.github.io/corona-india-status/covid19-indian-states.json",
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

$stateData = json_decode($response,true);
//looping and setting up statewise value
foreach ($stateData["data"] as $key => $value) {
  $keyRedis = strtolower($key);
  $keyRedis = str_replace(' ', '-', $keyRedis);
  foreach($value as $ink => $inv){
    `redis-cli hset $keyRedis $ink $inv `;
  }
}
echo "Done";

?>
