r.db("Corp")
.table("Stock")
.egJoin("IDDR", r.db("Corp").table("Plan"),{index:"IDRG"})
.zip()
.group("IDSTRUCT","IDDRG","NAME")
.ungroup()
.map({"Structure", r.row("group").nth(0),
      "Idrug", r.row("group").nth(1), 
      "Namedrug", r.row("group").nth(2),
      "SUM", r.row("reduction").sum("Ammount").default(0),
      "COUNT", r.row("reduction").sum("Quant").default(0),
      })
.filter(r.row("COUNT").lt(100)
.orderBy("COUNT")
