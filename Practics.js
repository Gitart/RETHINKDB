r.table("table").get("0").update(function(doc) {
  return doc.merge({shareRecords: doc("shareRecords").add(
     r.expr([{"update_Id":"update002"},{"update_Id":"update004"}]).filter(
        function(newRecord) {
           return doc("shareRecords").contains(
function(record) {
             return record("update_Id").eq(newRecord("update_Id"))
        }).not()
     })
   )}
)})



=> r.table("test").slice(5, 2, {:right_bound => "closed"})
irb(main):006:0> r.table('test')[5...2]
=> r.table("test").slice(5, 2, {:right_bound => "open"})


