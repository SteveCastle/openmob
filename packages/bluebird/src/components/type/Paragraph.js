import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  maxWidth,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledParagraph = styled.p`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${maxWidth}
  ${fontSize}
  ${color}
`;

function Paragraph({ children, dark }) {
  return (
    <StyledParagraph
      p={3}
      fontSize={3}
      width={1}
      maxWidth={6}
      color={dark ? 'type.light.f' : 'type.dark.f'}
    >
      {children}
    </StyledParagraph>
  );
}

Paragraph.propTypes = {
  children: PropTypes.element.isRequired,
  dark: PropTypes.bool
};
export default Paragraph;
