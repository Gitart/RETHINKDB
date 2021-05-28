
cls
echo START...
echo

rem curl -X POST -i http://ho.pharmbase.com.ua:5555/api/report/analysis/StockOnLine/ -d "{\"RELEVANCE\":\"24\",\"ANALYSIS_NAME\":\"StockOnLine\",\"IS_FAVORITE\":\"1\",\"ID_DRUGS\":[\"536214\",\"544170\",\"544172\",\"548190\"],\"ID_STRUCTURES\":[\"12379550\",\"12379551\",\"12557189\",\"12557166\",\"12557190\",\"12557167\",\"12557168\",\"12379556\"]}"  -H "AccessKey:49e65f74487338ce6f5d938084cac048efd40345" -H "AccessId:12379550"  -H "Content-Type:application/json" -H "Charset=utf-8" 
rem curl -X POST -i http://localhost:1277/add/ -d "{\"RELEVANCE\":\"24\",\"ANALYSIS_NAME\":\"StockOnLine\",\"IS_FAVORITE\":\"1\",\"ID_DRUGS\":[\"536214\",\"544170\",\"544172\",\"548190\"],\"ID_STRUCTURES\":[\"12379550\",\"12379551\",\"12557189\",\"12557166\",\"12557190\",\"12557167\",\"12557168\",\"12379556\"]}"  -H "AccessKey:49e65f74487338ce6f5d938084cac048efd40345" -H "AccessId:12379550"  -H "Content-Type:application/json" -H "Charset=utf-8"
REM curl -X POST -i http://localhost:1277/add/ -T m.json  -H  "AccessId:12379550"  -H "Content-Type:application/json" -H "Charset=utf-8"
curl -X POST -i http://localhost:1277/add/ -T "m1.json"  -H "Content-Type: application/json"  


pause
