import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap
} from 'styled-system';

const StyledMenu = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  overflow: hidden;
`;

function Menu({ children, vertical }) {
  return <StyledMenu>{children}</StyledMenu>;
}

Menu.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  vertical: PropTypes.bool
};

export default Menu;
