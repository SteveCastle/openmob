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
  height,
  alignItems,
  justifyContent,
  flexWrap,
  flexDirection
} from 'styled-system';

const StyledMenuBody = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${alignItems}
  ${justifyContent}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  ${flexDirection}
  overflow: scroll;
  position: fixed;
  top: 12%;
  left: 0;
`;

function MenuBody({ children }) {
  return (
    <StyledMenuBody
      flexDirection="column"
      fontSize={4}
      height={'88%'}
    >
      {children}
    </StyledMenuBody>
  );
}

MenuBody.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
};

export default MenuBody;
