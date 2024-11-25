export const convertConfig = (config: Record<string, any> , selectConvert:any) => {
  Object.keys(config).forEach((key) => {
    if (["education", "marital_status", "occupation"].includes(key)) {
      selectConvert[key] = config[key]
        ? config[key].split(",")
        : [];
    } else {
      selectConvert[key] = config[key];
    }
  });
};
