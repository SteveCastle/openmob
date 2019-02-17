import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledButton = styled.button`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  outline: none;
`;

function Button({ onClick, label, icon, loader, block, variant, state }) {
  return (
    <StyledButton
      block={block}
      border={0}
      borderRadius={2}
      onClick={onClick}
      p={3}
      m={1}
      bg={
        variant === 'dark'
          ? 'buttons.darkBackground'
          : 'buttons.lightBackground'
      }
      fontSize={1}
      color={variant === 'dark' ? 'buttons.lightText' : 'buttons.darkText'}
      width={block ? 1 : null}
    >
      {label}
    </StyledButton>
  );
}

Button.propTypes = {
  onClick: PropTypes.func,
  label: PropTypes.string,
  icon: PropTypes.oneOfType([PropTypes.element, PropTypes.string]),
  loader: PropTypes.oneOfType([PropTypes.element, PropTypes.string]),
  variant: PropTypes.oneOf(['primary', 'default', 'warning', 'dark'])
    .isRequired,
  state: PropTypes.oneOf(['ready', 'loading', 'success', 'error', 'disabled'])
    .isRequired,
  block: PropTypes.bool
};
Button.defaultProps = {
  state: 'ready'
};

export default Button;
