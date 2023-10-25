namespace go demoServer
struct TreeStructureResp {
    1: list<TreeStructureRespData> data;
    2: i32 code;
    3: string info;
}

struct TreeStructureRes{

}

struct TreeStructureRespData{
    1: string id;
    2: list<string> children;
    3: string name;
}

struct StuInfoRes {
    1: string id (api.query="id");
    2: string examId(api.query="examId");
}

struct StuInfoResp {
    1: i32 code;
    2: string info;
    3: StuInfoRespData data;
}
//belong to StuInfoResp
struct StuInfoRespData {
    1: string id;
    2: string name;
    3: string examId;   //Temporarily abandoned
    4: string examName;     //Temporarily abandoned
    5: list<KnowledgePointAccuracy> accuracy;
}

struct AllStuInfoRes {
}

struct AllStuInfoResp{
    1: i32 code;
    2: string info;
    3: list<StuInfoRespData> data
}


//belong to StuInfoRespData, ClassKnowledgeAccuracyPerResp
struct KnowledgePointAccuracy{
    1: string kid;
    2: string kpContent;
    3: double accuracy;
}

struct ExamListRes {

}

struct ExamListResp {
    1: i32 code;
    2: string info;
    3: list<ExamListRespData> data;
}

struct ExamListRespData{
    1: string id;
    2: string name;
}

// ClassKnowledgeCorrectPer
struct ClassKnowledgeCorrectPerRes {
    1: string classID (api.query="classID");
}

struct ClassKnowledgeAccuracyPerResp {
    1: i32 code;
    2: string info;
    3: string classID;
    4: list<KnowledgePointAccuracy> accuracy;
}

service DemoServer {
    TreeStructureResp  TreeStructure (TreeStructureRes res)(api.get="/treeStructure");
    StuInfoResp StuInformation (StuInfoRes res)(api.get="/stuInformation");
    AllStuInfoResp AllStuInformation (AllStuInfoRes res)(api.get="/allStuInfo");
    ExamListResp ExamList (ExamListRes res)(api.get="/examList");
    ClassKnowledgeAccuracyPerResp ClassKnowledgeCorrectPer(ClassKnowledgeCorrectPerRes res)(api.get="/classKnowledgeCorrectPer")
}