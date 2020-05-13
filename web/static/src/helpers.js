function ShouldRender(currProps, nextProps) {
  return nextProps.v !== currProps.v
}


export {ShouldRender};