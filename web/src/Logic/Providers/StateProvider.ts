import { createContext, useContext } from "react";
import AppStore from "../State/AppStore";
import { AuthStore } from "../State/AuthStore";
import { FeedBackStore } from "../State/FeedBackStore";

interface IStoresContext {
  appStore: AppStore;
  feedBackStore : FeedBackStore;
  authStore : AuthStore
}

export const StoresContext = createContext<IStoresContext>(
  {} as IStoresContext
);

export const useStores = () => useContext(StoresContext);
