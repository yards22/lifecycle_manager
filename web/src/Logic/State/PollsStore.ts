import { action, makeAutoObservable, observable } from "mobx";
import { PollsRepo } from "../Repository/PollsRepo";
import { MPolls } from "../Model/MPolls";

export class PollsStore{
    @observable isLoading:boolean = false
    @observable pollsArray : MPolls[]|null = []
    @observable token: string | null = null;
    pollsRepo:PollsRepo

    constructor(pollsRepo:PollsRepo){
        makeAutoObservable(this)
        this.pollsRepo = pollsRepo
        this.token = window.localStorage.getItem("token")
    }

    @action
    SetPollsArray = (x: MPolls[] | null) => {
      if (x !== null){
        this.pollsArray = x
      }
    };

    @action
    GetPolls = async ()=>{
        try {
            const {polls} = await this.pollsRepo.getPolls(this.token || "");
           this.pollsArray = polls
           console.log(polls);
           
        } catch (error) {
            throw error
        }
    }

    @action
    AddPoll = async (pollQuestion:string, pollOptions:string[], token:string)=>{
        try {
            const response = await this.pollsRepo.addPoll(pollQuestion ,pollOptions ,token);
            console.log("add poll response",response)
        } catch (error) {
            throw error
        }
    }

    
}