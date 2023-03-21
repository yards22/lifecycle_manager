export interface MFeedBack{
    feedback_id: number,
    user_id : number,
    image_uri : string | null,
    content : string | null,
    created_at : Date,
    status : boolean,
    comment : string |null

}