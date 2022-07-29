import { StyledButton } from "../components/StyledButton";
import { ComponentMeta } from "@storybook/react";

export default {
    title: 'StyledButton',
    component: StyledButton,
    argTypes: { onClick: {action: 'clicked'}}
} as ComponentMeta<typeof StyledButton>



export const Primary = (props) => {
    return (
        <StyledButton {...props} variant="primary">
            Primary
        </StyledButton>
    )
}

export const Success = (props) => {
    return (
        <StyledButton {...props} variant="success">
            Primary
        </StyledButton>
    )
}

export const Transparent = (props) => {
    return (
        <StyledButton {...props} variant="transparent">
            Transparent
        </StyledButton>
    )
}