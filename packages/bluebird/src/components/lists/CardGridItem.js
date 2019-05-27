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
import Card from '../cards/Card';

const StyledItem = styled.li`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${position}
  ${fontSize}
  ${color}
  box-sizing: border-box;
  cursor: pointer;
  transition: 2s;
  :hover {
    opacity: .7
  }
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

function GridItem({ width, title, caption, onClick }) {
  return (
    <StyledItem p={2} m={2} position="relative" width={width} onClick={onClick}>
      <Card>
        <StyledTitle color="gray.9" m={4}>
          {title}
        </StyledTitle>
        <StyledCaption color="gray.9" mx={4} my={3}>
          {caption}
        </StyledCaption>
      </Card>
    </StyledItem>
  );
}

GridItem.propTypes = {
  width: PropTypes.number,
  title: PropTypes.string,
  caption: PropTypes.string,
  onClick: PropTypes.func,
};
GridItem.defaultProps = { width: 1 / 4 };

export default GridItem;
