import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  position,
  bottom,
  borderRadius
} from 'styled-system';

const StyledItem = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${position}
  ${fontSize}
  ${color}
`;

const StyledImage = styled.img`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;
const StyledTitle = styled.h2`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
`;

function GridItem({ uri, width, height, title, id, slug }) {
  return (
    <StyledItem  m={2} position="relative">
      <StyledImage src={uri} width={1} m={0} p={0}/>
      <StyledTitle color="gray.0" position="absolute" m={4} bottom="0">{title}</StyledTitle>
    </StyledItem>
  );
}

GridItem.propTypes = {
  uri: PropTypes.string,
  width: PropTypes.number,
  height: PropTypes.number,
  title: PropTypes.string,
  id: PropTypes.string,
  slug: PropTypes.string
};
GridItem.defaultProps = { size: 12 };

export default GridItem;
