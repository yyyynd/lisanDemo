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
//belong to StuInfoRespData
struct KnowledgePointAccuracy{
    1: string kid;
    2: double accuracy;
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

service DemoServer {
    TreeStructureResp  TreeStructure (TreeStructureRes res)(api.get="/treeStructure");
    StuInfoResp StuInformation (StuInfoRes res)(api.get="/stuInformation");
    ExamListResp ExamList (ExamListRes res)(api.get="/examList");
}