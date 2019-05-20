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
  borderRadius,
} from 'styled-system';

const StyledItem = styled.li`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${position}
  ${fontSize}
  ${color}
  box-sizing: border-box;
  transition: 2s;
  :hover {
    opacity: .7
  }
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

const StyledCaption = styled.p`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
  text-shadow: 2px 4px 3px rgba(0,0,0,0.3);
`;

function GridItem({ uri, width, height, title, id, slug, caption }) {
  return (
    <StyledItem p={2} position="relative" width={width}>
      <StyledImage src={uri} width={1} m={0} p={0} />
      <StyledTitle color="gray.0" position="absolute" m={4} bottom="0">
        {title}
      </StyledTitle>
      <StyledCaption
        color="gray.0"
        position="absolute"
        mx={4}
        my={3}
        bottom="0"
      >
        {caption}
      </StyledCaption>
    </StyledItem>
  );
}

GridItem.propTypes = {
  uri: PropTypes.string,
  width: PropTypes.number,
  height: PropTypes.number,
  title: PropTypes.string,
  caption: PropTypes.string,
  id: PropTypes.string,
  slug: PropTypes.string,
};
GridItem.defaultProps = { width: 1 / 4 };

export default GridItem;
