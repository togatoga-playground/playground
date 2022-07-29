import { NextPage } from "next";
import styled from "styled-components";

const Text = styled.span`
  color: ${(props) => props.theme.colors.primary};
  font-size: ${(props) => props.theme.fontSizes[0]};
  margin: ${(props) => props.theme.space[2]};
`;
const Page: NextPage = () => {
  return (
    <div>
      <Text>Themeから参照した色を使用しています</Text>
    </div>
  );
};

export default Page;