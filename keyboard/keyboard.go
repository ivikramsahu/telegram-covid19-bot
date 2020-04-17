package keyboard

import(
  "github.com/go-telegram-bot-api/telegram-bot-api"
)

//AsiaKeyboard countries keyboard
var AsiaKeyboard =`
-----------------------------------------
            Asian countries               
-----------------------------------------      
1. /india           25./cambodia
2. /china           26./jordan
3. /indonesia       27./azerbaijan
4. /pakistan        28./uae
5. /bangalesh       29./tajikistan
6. /japan           30./israel
7. /philippines     31./laos
8. /vietnam         32./lebanon
9. /turkey          33./kyrgyzstan
10./iran            34./turkmenistan
11./thailand        35./singapore
12./myanmar         36./oman
13./south_korea     37./stateofpalestine
14./iraq            38./kuwait
15./afghanistan     39./georgia
16./saudi_arabia    40./mongolia
17./uzbekistan      41./armenia
18./malaysia        42./qatar
19./yemen           43./bahrain
20./nepal           44./timorleste
21./north_korea     45./cyprus
22./srilanka        46./bhutan
23./kazakhstan      47./maldives
24./syria           48./brunei
`
//AfricaKeyboard continents keyboard
var AfricaKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("East Africa"),
    tgbotapi.NewKeyboardButton("West Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("North Africa"),
    tgbotapi.NewKeyboardButton("South Africa"),
  ),
)

//EuropeKeyboard countries keyboard
var EuropeKeyboard = `
-----------------------------------------
European countries               
-----------------------------------------      
1. /russia          23./finland
2. /germany         24./slovakia
3. /UK              25./norway
4. /france          26./ireland
5. /italy           27./croatia
6. /spain           28./moldova
7. /ukraine         29./bosnia_and_herzegovina
8. /poland          30./albania
9. /romania         31./lithuania
10./netherlands     32./north_macedonia
11./belgium         33./slovenia
12./czech_republic  34./latvia
13./greece          35./estonia
14./portugal        36./montenegro
15./sweden          37./luxembourg
16./hungary         38./malta
17./belarus         39./andorra
18./austria         40./iceland
19./serbia          41./monaco
20./switzerland     42./liechtenstein
21./bulgaria        43./san_marino
22./denmark         44./holy_see
-----------------------------------------
`

//AmericaKeyboard countries keyboard
var AmericaKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Asia"),
    tgbotapi.NewKeyboardButton("Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("America"),
    tgbotapi.NewKeyboardButton("Europe"),
  ),
)
