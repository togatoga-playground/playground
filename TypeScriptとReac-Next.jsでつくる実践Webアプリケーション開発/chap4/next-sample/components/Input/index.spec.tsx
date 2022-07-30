import {render, screen, RenderResult, fireEvent} from '@testing-library/react'
import { Input } from '.'
describe('Input', () => {
    let renderResult: RenderResult
    beforeEach(() => {
        renderResult = render(<Input id="username" label="Username"/>)
    }) 

    afterEach(() => {
        renderResult.unmount()
        
    }) 
    it('should empty in input on initial render', () => {
        const inputNode = screen.getByLabelText('Username') as HTMLInputElement
        expect(inputNode).toHaveValue('')
    })

    it('should show input text', () => {
        const inputText = 'Test Input Text'
        const inputNode = screen.getByLabelText('Username') as HTMLInputElement
        fireEvent.change(inputNode, {target: {value: inputText}})

        expect(inputNode).toHaveValue(inputText)
    })

    it('should reset when user clicks button', async() => {
        const inputTeext = 'Test Input Text'
        const inputNode = screen.getByLabelText('Username') as HTMLInputElement
    })
})