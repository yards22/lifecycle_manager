import { Observer } from "mobx-react-lite";
import styled from "styled-components";
import { useStores } from "../../Logic/Providers/StateProvider";
import FeedBackCard from "./FeedBackCard";
import { useEffect } from "react";

const SFeedBackScreenIndex = styled.div`
  display: grid;
  grid-template-columns: ${(p) =>
    p.theme.deviceWidth > 1600
      ? "auto auto auto"
      : p.theme.deviceWidth < 800
      ? "auto"
      : "auto auto"};
  width: 100%;
  flex-wrap: wrap;
  /* border: 1px solid ; */
`;

function FeedBackScreenIndex() {
  const stores = useStores();

  console.log("here at feedback screen")

  useEffect(() => {
    stores.feedBackStore.GetFeedBacks();
  }, []);

  const { feedBackStore, appStore } = stores;
  return (
    <SFeedBackScreenIndex theme={{ deviceWidth: appStore.deviceWidth }}>
      <Observer>
        {() => {
          const { feedbackArray } = feedBackStore;
          if (feedbackArray.length === 0) return <p>No feedback</p>;
          return (
            <div>
              {feedbackArray.map((item, index) => {
                return (
                  <FeedBackCard feedBack={item} key={`normal_post_${index}`} />
                );
              })}
            </div>
          );
        }}
      </Observer>
    </SFeedBackScreenIndex>
  );
}

export default FeedBackScreenIndex;
