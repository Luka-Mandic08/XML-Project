# Generate new feature

## Step 1: Generation of feature folder

Before generation of new feature first thing one needs to think about is whether it would be used only in fights app or it may be shared between apps( currently only flights, but in future there could be more apps added, e.g. airbnb). After that decision, one needs to create feature folder in one of two folders under [features library](../libs/features/).

- feature that will be used only in flights app should go to the [flights folder](../libs/features/flights)
- feature that will be shared between apps should go to the [shared folder](../libs/features/shared)

Example: we are developing new feature for handling reception and assume this will only be used in flights app, therefore first we will do is to create `reception` folder under [flights features folder](../libs/features/flights).
Once feature folder is created next step is to create feature specific libraries.

## Step 2: Generation of feature specific libraries

Once you have understood which types of libraries do we have and what is purpose of each of them, you should create appropriate libraries. Let's get back to example mentioned above, and assume, we are creating `reception` feature, where we will have a list of receptions items displayed. For loading receptions, we need to call backend api. So, let's say that we will need:

- service to call backend api(e.g. receptions-service)
- one component, that will be our container and contain encapsulated data; it will call service method for loading data and render list of Reception components (e.g. ReceptionsContainer component)
- model

From this scenario, we assume that we need to create containers library(for ReceptionsContainer) and data-access folder with services(for projects-service).

In next steps there is detailed explanation for creation of libraries and components and it is assumed that one has already installed Nx Console extension.

### Step 2.1. Creation of library

1. Easiest way to create new library is right clicking on folder where library should be placed (in this case it would be previously created `reception` folder) and then selecting option Nx generate which should open menu on the top of workspace

   ![Choose Nx generate](Images/Frontend_HowToCreateNewFeature/1.png)

2. From opened menu choose option `@nrwl/react-library` (tip: type in 'react-library' into autocomplete to easier find it).

   ![Choose @nrwl/react-library](Images/Frontend_HowToCreateNewFeature/2.png)

3. Now, enter library name(for container components we use lowercase and add suffix `-container`, the name of functional component (for container) itself should be changed to PascalCase `ComponentContainer`), we use `container`, for data-acces we use `data-access`, and uncheck `Generate a default component` checkbox.

   ![Set library name and uncheck Generate default component checkbox](Images/Frontend_HowToCreateNewFeature/3.png)

4. Click `Run` button from upper right corner and then you should see library generated in projects tree. Now, you are ready to add containers inside containers/src folder.

   ![Generated library in projects tree](Images/Frontend_HowToCreateNewFeature/5.png)

### Step 2.4. Creation of models

All models are created inside `models` library. To create new model, create new folder inside `models/src/lib`. Inside new folder create `.ts` files with your models. Make sure to export your models inside `models/src/index.ts` file.

![Choose unitTestRunner none](Images/Frontend_HowToCreateNewFeature/7.png)

## Step 3: Generation of new component

1. Easiest way to create new component is, same as for creating new library, right click on folder where component should be placed(that is `src` or `src/lib` folder of some container) and selection of option Nx generate.
2. From menu choose option `@nrwl/react-component` (tip: type 'react-component' into autocomplete to easier find it).
3. Enter component name (e.g. we want to create ProjectsContainer component we set name to projects-container as it is going to be file name and we use dash-case here; component name will be in PascalCase by default) and check `export` checkbox (in this way we assume component will be exported from lib's index.ts file and therfore could be imported in other components)

   ![Enter component name and check export checkbox](Images/Frontend_HowToCreateNewFeature/8.png)

4. Click `Run` button from upper right corner and find your component under src/lib folder.

   > NOTE: Component will be generated inside projects-container folder together with projects-container.spec.tsx file for writing tests and projects-container.module.scss file for writing styles.

   > NOTE: When creating models and services, you can do it "by hand" (create appropraite `.ts` files and add export for this files to `index.ts`). Another option is to use Nx generate in the way described above for generating components with some differencies. One need to set `style` option to `none`, as style files are not needed neither for services nor for models. Also, when creating models, one needs to check `skipTests` checkbox and `flat` checkbox(unless there is real need to create additional folder for models organization). Do not forget, Nx will generate `.tsx` files, but then they need to be changed to `.ts` files.
