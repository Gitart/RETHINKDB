joker_pairs = function (user, bucket) {
  return(db.table("jokes").getAll(bucket, {index: "bucket"}).map(function (r) {
    return([user, r("user")]);
  }).coerceTo("array"));
};
 
mutual_jokes = function (user, bucket) {
  return(db["do"](joker_pairs(user, bucket), function (pairs) {
    return(db.branch(pairs.count().gt(0), db.table("jokes").getAll(db.args(pairs), {index: "for"}).count().sub(1), 0));
  }));
};
 
joker_stats = function (r) {
  var user = r("user");
  var bucket = r("bucket")("id");
  return({mutual: mutual_jokes(user, bucket)});
};
 
joke_query = function (user, bucket, max, stats) {
  var i = db.table("buckets").get(db.row("bucket"));
  var u = db.table("users").get(db.row("jokee"));
  var q = joke_pool(user, bucket).merge({bucket: i}).merge({jokee: u});
  if (!bucket) {
    q = q.limit(max || 50);
  }
  if (stats) {
    q = q.merge(joker_stats);
  }
  if (bucket) {
    q = q.without("bucket").without("jokee");
  }
  return(q);
};
 
var profile = [
  {
    "description": "Evaluating without.",
    "duration(ms)": 20295.45191,
    "sub_tasks": [
      {
        "description": "Evaluating without.",
        "duration(ms)": 20292.697316,
        "sub_tasks": [
          {
            "description": "Evaluating merge.",
            "duration(ms)": 20289.473542,
            "sub_tasks": [
              {
                "description": "Evaluating merge.",
                "duration(ms)": 89.063321,
                "sub_tasks": [
                  {
                    "description": "Evaluating merge.",
                    "duration(ms)": 66.476687,
                    "sub_tasks": [
                      {
                        "description": "Evaluating orderby.",
                        "duration(ms)": 38.144123,
                        "sub_tasks": [
                          {
                            "description": "Evaluating get_all.",
                            "duration(ms)": 0.041009,
                            "sub_tasks": [
                              {
                                "description": "Evaluating table.",
                                "duration(ms)": 0.022384,
                                "sub_tasks": [
                                  {
                                    "description": "Evaluating datum.",
                                    "duration(ms)": 0.000676,
                                    "sub_tasks": []
                                  },
                                  {
                                    "description": "Evaluating db.",
                                    "duration(ms)": 0.008283,
                                    "sub_tasks": [
                                      {
                                        "description": "Evaluating datum.",
                                        "duration(ms)": 0.000081,
                                        "sub_tasks": []
                                      }
                                    ]
                                  }
                                ]
                              },
                              {
                                "description": "Evaluating datum.",
                                "duration(ms)": 0.00014,
                                "sub_tasks": []
                              },
                              {
                                "description": "Evaluating datum.",
                                "duration(ms)": 0.000077,
                                "sub_tasks": []
                              }
                            ]
                          },
                          {
                            "description": "Evaluating desc.",
                            "duration(ms)": 0.000775,
                            "sub_tasks": [
                              {
                                "description": "Evaluating datum.",
                                "duration(ms)": 0.000086,
                                "sub_tasks": []
                              }
                            ]
                          },
                          {
                            "description": "Perform read.",
                            "duration(ms)": 2.158814,
                            "sub_tasks": [
                              {
                                "parallel_tasks": [
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.628261,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.502125,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.952707,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.796865,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.857791,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.708561,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.985939,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.835986,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.615992,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.51066,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 0.8066,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.691181,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 1.5849,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 1.08066,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ],
                                  [
                                    {
                                      "description": "Perform read on shard.",
                                      "duration(ms)": 1.009471,
                                      "sub_tasks": [
                                        {
                                          "description": "Do range scan on secondary index.",
                                          "duration(ms)": 0.855559,
                                          "sub_tasks": []
                                        }
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          },
                          {
                            "description": "Sorting in-memory.",
                            "mean_duration(ms)": 0.00802,
                            "n_samples": 994
                          }
                        ]
                      },
                      {
                        "description": "Evaluating stream eagerly.",
                        "mean_duration(ms)": 0.173564,
                        "n_samples": 161
                      }
                    ]
                  },
                  {
                    "description": "Evaluating stream eagerly.",
                    "mean_duration(ms)": 0.137301,
                    "n_samples": 161
                  }
                ]
              },
              {
                "description": "Evaluating stream eagerly.",
                "mean_duration(ms)": 125.142485,
                "n_samples": 161
              }
            ]
          },
          {
            "description": "Evaluating stream eagerly.",
            "mean_duration(ms)": 0.017751,
            "n_samples": 161
          }
        ]
      },
      {
        "description": "Evaluating stream eagerly.",
        "mean_duration(ms)": 0.015904,
        "n_samples": 161
      }
    ]
  }
];
