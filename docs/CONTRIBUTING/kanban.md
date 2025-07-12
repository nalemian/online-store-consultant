### Kanban board
The link to our Kanban board:
https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/boards

#### Columns
There are entry criteria for each column. An issue can be closed when it reaches the ```Done``` column.

#### To Do
```
[Entry Criteria]
* The issue is unambigiously formulated as the issue form template.
* The issue is unanimously estimated in story points by all team members.
* The label is attached to the issue.
* The issue is desided to be in the sprints.
```

#### In Progress
```
[Entry Criteria]
* The issue description was revised to provide missing details.
* The issue was added to the current sprint.
```

#### Ready to Deploy
```
[Entry criteria]
* The MR attached to the issues passes the pipeline stage.
* The MR is approaved by at least one team member.
* All acceptance criterias are satisfied.
* All issues on which it depends are done.
```

#### Closed/Done
```
[Entry criteria]
* Also the previous development stages are passed.
* The code is merged into the main branch.
```