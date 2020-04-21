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

for ($i=0;$i<=sizeof($cityData);$i++){
  foreach($cityData[$i] as $k => $v){
    echo "key :: $k --> value :: $v\n";
    if ($k == "districtData"){
      var_dump($k);
      exit;
      foreach($k as $ink => $inv){
        echo "key1 :: $ink ---> Value :: $inv\n";
        exit;   
      }
    }
  }
}
exit;
var_dump($cityData[1]);
echo "Done";

?>
