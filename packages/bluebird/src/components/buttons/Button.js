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

function Button({ onClick, label, icon, loader, block, dark, state }) {
  return (
    <StyledButton
      block={block}
      border={0}
      borderRadius={2}
      onClick={onClick}
      p={3}
      bg={dark ? 'buttons.darkBackground' : 'buttons.lightBackground'}
      fontSize={1}
      color={dark ? 'buttons.lightText' : 'buttons.darkText'}
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
  state: PropTypes.oneOf(['ready', 'loading', 'success', 'error']).isRequired,
  block: PropTypes.bool,
  dark: PropTypes.bool
};
Button.defaultProps = {
  state: 'ready'
};

export default Button;
