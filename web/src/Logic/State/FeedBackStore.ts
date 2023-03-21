import { action, makeAutoObservable, observable } from "mobx";
import { MFeedBack } from "../Modal/MFeedBack";
import { FeedBackRepo } from "../Repository/FeedBackRepo";

export class FeedBackStore{
    @observable isLoading : boolean = false;
    @observable feedbackArray : MFeedBack[] = [];
    @observable token:string|null =null
    feedBackRepo : FeedBackRepo

    constructor(feedBackRepo:FeedBackRepo){
        makeAutoObservable(this)
        this.feedBackRepo = feedBackRepo
        this.token = window.localStorage.getItem("token");
    }

    @action
    SetFeedbackArray = (x: MFeedBack[] | null) => {
      if (x !== null){
        this.feedbackArray = x
      }
    };


    @action
    SetIsLoading = (loading : boolean) =>{
        this.isLoading = loading
    }

    @action
    GetFeedBacks = async ()=> {
    try {
      const feedback:any = await this.feedBackRepo.getFeedBacks( this.token || "");
      console.log(feedback)
      this.SetFeedbackArray(feedback.feedback);
    } catch (err) {
      throw err;
    }
  }

  @action
  PostFeedBackComments = async(status : boolean, comment:string | null,idx :number)=>{
    try{
      const feedbacks = this.feedbackArray.map((v) => v);
      console.log(feedbacks)
      const res = await this.feedBackRepo.postFeedbackComments(status,comment,idx,this.token || "")
      feedbacks[idx-2].comment  = comment;
      feedbacks[idx-2].status = status;
    }
    catch(err){
      throw err;
    }
   }
}
