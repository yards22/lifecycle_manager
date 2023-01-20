import { Carousel } from "@mantine/carousel";
import styled from "styled-components";
import { useStores } from '../../Logic/Providers/StateProvider'
interface NormalPostMediaProps {
  media: string | null;
}

const SFeedPostImage = styled.img`
  width: 100%;
  object-fit: contain;
  max-height: 400px;
`;

function FeedBackMedia(props: NormalPostMediaProps) {
  const stores = useStores();
  console.log("in feedback media",props.media)
  return (
    <div
      style={{
        margin: "10px 0",
      }}
    >{
     props.media !== null ? <SFeedPostImage src={props.media} /> : null
    }
    </div>
  );
}

export default FeedBackMedia;
