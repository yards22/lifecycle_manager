import { Observer } from 'mobx-react-lite'
import styled from 'styled-components'
import { DummyTestDataFeedBack } from '../../Data/DummyFeedBackData'
import { useStores } from '../../Logic/Providers/StateProvider'
import FeedBackCard from './FeedBackCard'

const SFeedBackScreenIndex = styled.div`
   display: grid;
   grid-template-columns: ${(p)=> p.theme.deviceWidth > 1600 ? "auto auto auto": p.theme.deviceWidth < 800 ? "auto":"auto auto"};
   width: 100%;
   flex-wrap: wrap;
   /* border: 1px solid ; */
`

function FeedBackScreenIndex() {
  const stores = useStores();
  return (
    <Observer>
      {
        ()=>{
          const {appStore} = stores
          return(
            <SFeedBackScreenIndex theme={{deviceWidth : appStore.deviceWidth}}>
                {
                  DummyTestDataFeedBack.map((each)=>{
                    return(
                      <FeedBackCard key={"feedbackNo"+each.feedback_id} feedBack={each} />
                    )
                  })
                }
            </SFeedBackScreenIndex>
          )
        }
      }
    </Observer>
  )
}

export default FeedBackScreenIndex