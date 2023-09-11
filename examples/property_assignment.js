db.t_user_profile.aggregate([
    {$match:{ "tag_TI836670048880824320_TN836671859905794048":{$exists:true},"tag_TI579011270012837888_TN579011444302946304":{$exists:false} }},
    {$limit:10}
    ]);