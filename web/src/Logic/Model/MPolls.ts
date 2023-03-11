export interface MPolls{
    poll_by:  String,
    poll_question: String,
    options_count:number,
    options : String,
    created_at : Date

} 

export interface MPollButton{
    title : String,
    votes : number
}