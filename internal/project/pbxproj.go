package project

var pbxprojTemplate = parseTemplate(`// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 48;
	objects = {

/* Begin PBXGroup section */
		0B9773641D8AE37900017700 /* Products */ = {
			isa = PBXGroup;
			children = (
			);
			name = Products;
			sourceTree = "<group>";
		};
		0BA02E491D87571500F1E8D3 = {
			isa = PBXGroup;
			children = (
				0B9773641D8AE37900017700 /* Products */,
			);
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXProject section */
		0BA02E4A1D87571500F1E8D3 /* Project object */ = {
			isa = PBXProject;
			attributes = {
				LastUpgradeCheck = 0800;
			};
			buildConfigurationList = 0BA02E4D1D87571500F1E8D3 /* Build configuration list for PBXProject "{{.ProjectName}}" */;
			compatibilityVersion = "Xcode 8.0";
			developmentRegion = English;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
			);
			mainGroup = 0BA02E491D87571500F1E8D3;
			productRefGroup = 0B9773641D8AE37900017700 /* Products */;
			projectDirPath = "";
			projectRoot = "";
			targets = (
			);
		};
/* End PBXProject section */

/* Begin XCBuildConfiguration section */
		0BA02E4E1D87571500F1E8D3 /* Debug */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
			};
			name = Debug;
		};
		0BA02E4F1D87571500F1E8D3 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
			};
			name = Release;
		};
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		0BA02E4D1D87571500F1E8D3 /* Build configuration list for PBXProject "{{.ProjectName}}" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				0BA02E4E1D87571500F1E8D3 /* Debug */,
				0BA02E4F1D87571500F1E8D3 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = 0BA02E4A1D87571500F1E8D3 /* Project object */;
}
`)
