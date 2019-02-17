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

const StyledParagraph = styled.p`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Paragraph({ children, dark }) {
  return (
    <StyledParagraph
      p={3}
      fontSize={3}
      width={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
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
