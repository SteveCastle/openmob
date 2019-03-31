import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  fontFamily,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledSubHeader = styled.h4`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${fontFamily}
  ${color}
`;

function SubHeader({ children, dark }) {
  return (
    <StyledSubHeader
      p={3}
      pt={0}
      mt={0}
      pb={0}
      mb={0}
      fontSize={4}
      fontFamily="sansSerif"
      width={1}
      color={dark ? 'type.light.subHeader' : 'type.dark.subHeader'}
    >
      {children}
    </StyledSubHeader>
  );
}

SubHeader.propTypes = {
  children: PropTypes.element.isRequired,
  dark: PropTypes.bool
};
export default SubHeader;
