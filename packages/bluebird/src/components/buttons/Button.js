import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  themeGet,
  space,
  width,
  fontSize,
  color,
  borders,
  boxShadow,
  borderRadius
} from 'styled-system';

const StyledButton = styled.button`
  ${borders}
  ${boxShadow}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  outline: none;
  box-sizing: border-box;
  cursor: pointer;

  &:hover{
    background-color: ${themeGet('colors.buttons.light.primary.hover')};
  }
  &:active{
    box-shadow: none;
  }
`;

function Button({
  onClick,
  label,
  icon,
  loader,
  block,
  variant = 'default',
  state,
  dark
}) {
  return (
    <StyledButton
      block={block}
      border={0}
      boxShadow={0}
      borderRadius={2}
      onClick={onClick}
      p={3}
      bg={dark ? `buttons.dark.${variant}.bg` : `buttons.light.${variant}.bg`}
      fontSize={1}
      color={
        dark
          ? `buttons.dark.${variant}.label`
          : `buttons.light.${variant}.label`
      }
      width={block ? 1 : null}
    >
      {label}
    </StyledButton>
  );
}

Button.propTypes = {
  onClick: PropTypes.func,
  label: PropTypes.string,
  icon: PropTypes.oneOfType([
    PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
    PropTypes.string
  ]),
  loader: PropTypes.oneOfType([
    PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
    PropTypes.string
  ]),
  variant: PropTypes.oneOf(['primary', 'default', 'warning', 'outline'])
    .isRequired,
  state: PropTypes.oneOf(['ready', 'loading', 'success', 'error', 'disabled'])
    .isRequired,
  block: PropTypes.bool,
  dark: PropTypes.bool
};
Button.defaultProps = {
  state: 'ready'
};

export default Button;
