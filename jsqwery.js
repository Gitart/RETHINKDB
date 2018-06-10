r.db("videojet").table("testdm").insert(
  r.db("videojet")
    .table("data_value")
    .getAll("687D2CD7-4B0A-4EE4-A4F6-2F67C0CDA449", {index: "printerid"})
    .innerJoin(
      r.db("videojet").table("model_msg_data").getAll("65605151",{index: "modelmsgdataid"}),
      function (datavalue, modelmsg) {
        return datavalue("modelmsgdataid").eq(modelmsg("modelmsgdataid"));
      }
    )
    .map({
      printerid: r.row("left")("printerid"),
      modelmsgdataid: r.row("left")("modelmsgdataid"),
      datetimestamp: r.row("left")("datetimestamp"), 
      datatagname: r.row("right")("datatagname")
    })
    .innerJoin(
      r.db("videojet").table("printer"),
      function (datavalue,printer) {
        return datavalue("printerid").eq(printer("printerid"));
      }
    )
    .map({
      printerid: r.row("left")("printerid"),
      modelmsgdataid: r.row("left")("modelmsgdataid"),
      datetimestamp: r.row("left")("datetimestamp"),
      datatagname: r.row("left")("datatagname"),
      modelid: r.row("right")("modelid"),
      serialnumber: r.row("right")("serialnumber")
    })
    .innerJoin(
      r.db("videojet").table("model"),
      function (datavalue,model) {
        return datavalue("modelid").eq(model("modelid"));
      }
    )
    .map({
      printerid: r.row("left")("printerid"),
      modelmsgdataid: r.row("left")("modelmsgdataid"),
      datetimestamp: r.row("left")("datetimestamp"),
      datatagname: r.row("left")("datatagname"),
      modelid: r.row("left")("modelid"),
      serialnumber: r.row("left")("serialnumber"),
      modelname: r.row("right")("modelname")
    })
 )
