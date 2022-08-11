/* eslint-disable @typescript-eslint/ban-types */
import { theme } from '../themes'

export type AppTheme = typeof theme

type SpaceThemeKeys = keyof typeof theme.space
type ColorThemeKeys = keyof typeof theme.colors
type FontSizeThemeKeys = keyof typeof theme.fontSizes
type LetterSpacingThemeKeys = keyof typeof theme.letterSpacings
type LineHeightThemeKeys = keyof typeof theme.lineHeights

export type Space = SpaceThemeKeys | (string & {})
export type Color = ColorThemeKeys | (string & {})
export type FontSize = FontSizeThemeKeys | (string & {})
export type LetterSpacing = LetterSpacingThemeKeys | (string & {})
export type LineHeight = LineHeightThemeKeys | (string & {})

const BREAKPOINTS: { [key: string]: string } = {
    sm: '640px',
    md: '768px',
    lg: '1024px',
    xl: '1280px',
}

export function toPropValue<T>(
    propKey: string,
    prop?: Responsive<T>,
    theme?: AppTheme,
) {
    if (prop === undefined) return undefined
    if ()
}