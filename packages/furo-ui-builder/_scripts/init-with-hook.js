#!/usr/bin/env node

const fs = require('fs');
const execSync = require('child_process').execSync;
const path = require('path');
const U33eBuilder = require("./u33eBuilder");
const Helper = require("./InitHelper");
let config;
// config öffnen
if (fs.existsSync('./furo.ui.spec.conf.json')) {
  config = JSON.parse(fs.readFileSync('./furo.ui.spec.conf.json'));
} else {
  console.log("furo.ui.spec.conf.json not found, you can copy an example from " + path.normalize(__dirname + "/../"));
  process.exit(1);
}

function sh(command, arguments) {
  execSync(command + " " + arguments.join(" "), {stdio: 'inherit'});
}

const SpecDir = config.spec_dir;
const UiSpecDir = config.ui_spec_out;

sh("mkdir -p", [UiSpecDir]);

// prepare hooks
const hooks = {type: [], service: []};
// type hooks
config.hooks.type.forEach((hook) => {
  hooks.type.push(require(process.cwd() + "/" + hook));
});
// service hooks
config.hooks.service.forEach((hook) => {
  hooks.service.push(require(process.cwd() + "/" + hook));
});


let speclist = Helper.walkSync(SpecDir).filter((filepath) => {
  let filename = path.basename(filepath);
  if (filename.indexOf("_") >= 0) {
    // ignore type_collection or type_entity
    return false;
  }

  // skip spec files
  if (config.skip_spec.indexOf(filename) != -1) {
    console.log("skip:", filename);
    return false;
  }
  // return all specs
  return (filename.indexOf(".spec") > 0);

});


// resolve all possible created names and paths for the import helper
speclist.forEach((pathToSpec) => {
  let ctx = Helper.specInfo(pathToSpec);
  ctx.config = config;
  // load spec file
  ctx.spec = JSON.parse(fs.readFileSync(pathToSpec));
  ctx.package = ctx.spec.__proto.package;

  // loop hooks for service or type
  hooks[ctx.kindOf].forEach((hook) => {
    let ctx = Helper.specInfo(pathToSpec);
    ctx.config = config;
    // load spec file
    ctx.spec = JSON.parse(fs.readFileSync(pathToSpec));
    ctx.package = ctx.spec.__proto.package;

    hook.getPath(ctx);

  });
});


speclist.forEach((pathToSpec) => {
  let ctx = Helper.specInfo(pathToSpec);
  ctx.config = config;
  // load spec file
  ctx.spec = JSON.parse(fs.readFileSync(pathToSpec));
  ctx.package = ctx.spec.__proto.package;

  // loop hooks for service or type
  hooks[ctx.kindOf].forEach((hook) => {
    let u33e = new hook(ctx, new U33eBuilder());

    if (u33e instanceof U33eBuilder) {
      sh("mkdir -p", [path.dirname(u33e.model.path)]);
      // write u33e file if model is returned
      fs.writeFileSync(u33e.model.path, u33e.getU33e());

    }
  });
});
