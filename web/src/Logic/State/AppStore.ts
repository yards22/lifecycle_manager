import { observable, action, makeAutoObservable } from "mobx";

export default class AppStore {
  @observable currentTab: number = 0;
  @observable deviceWidth: number = 0;
  @observable selectedDate: Date = new Date();
  @observable isPhone: boolean = false;
  @observable isTablet: boolean = false;
  @observable isDesktop: boolean = true;
  @observable navigationState: number = 0;
  @observable currentWindowSize: number = window.innerWidth;

  constructor() {
    makeAutoObservable(this);
  }

  @action
  setIsPhone = (isPhone: boolean) => {
    this.isPhone = isPhone;
    this.isTablet = false;
    this.isDesktop = false;
  };

  @action
  setIsTablet = (isTablet: boolean) => {
    this.isTablet = isTablet;
    this.isPhone = false;
    this.isDesktop = false;
  };

  @action
  setIsDesktop = (isDesktop: boolean) => {
    this.isDesktop = isDesktop;
    this.isPhone = false;
    this.isTablet = false;
  };

  @action
  SetCurrentTab = (currentTab: number) => {
    this.currentTab = currentTab;
  };

  @action
  SetCurrentDate = (date: Date) => {
    this.selectedDate = date;
  };

  @action
  setDeviceWidth = (state: number) => {
    this.deviceWidth = state;
  };

  @action
  setNavigationState = (state: number) => {
    this.navigationState = state;
  };
}
