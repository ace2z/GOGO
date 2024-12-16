#!/bin/bash
export TPUTBIN=`which tput`
$TPUTBIN setaf 5

VERSION="v7.03"

export GO111MODULE=on
export GOPROXY=direct
#export GOSUMDB=off


ECHO_COLOR() {
  echo -n "$($TPUTBIN setaf ${1})${2}$($TPUTBIN sgr0)"
  if [[ -z ${3} ]]; then
    echo ""
  fi
}

pressany() {
  read -rs -p"Press any key to continue"; echo
}

# # Useful aliases (maybe?)
# red=1
# green=2
# yellow=3
# blue=4
# magenta=5
# cyan=6
# grey=7
# gray=7
# white=7
# alias echo_color="ECHO_COLOR ${1} ${2} ${3}"
# alias echo_green="echo_color 2 ${1}"
# alias echo_yellow="echo_color 3 ${1}"
# alias echocolor="echo_color"

# Build mode reference:   https://pkg.go.dev/cmd/go
# -buildmode=archive
# 	Build the listed non-main packages into .a files. Packages named
# 	main are ignored.

# -buildmode=c-archive
# 	Build the listed main package, plus all packages it imports,
# 	into a C archive file. The only callable symbols will be those
# 	functions exported using a cgo //export comment. Requires
# 	exactly one main package to be listed.

# -buildmode=c-shared
# 	Build the listed main package, plus all packages it imports,
# 	into a C shared library. The only callable symbols will
# 	be those functions exported using a cgo //export comment.
# 	Requires exactly one main package to be listed.

# -buildmode=default
# 	Listed main packages are built into executables and listed
# 	non-main packages are built into .a files (the default
# 	behavior).

# -buildmode=shared
# 	Combine all the listed non-main packages into a single shared
# 	library that will be used when building with the -linkshared
# 	option. Packages named main are ignored.

# -buildmode=exe
# 	Build the listed main packages and everything they import into
# 	executables. Packages not named main are ignored.

# -buildmode=pie
# 	Build the listed main packages and everything they import into
# 	position independent executables (PIE). Packages not named
# 	main are ignored.

# -buildmode=plugin
# 	Build the listed main packages, plus all packages that they
# 	import, into a Go plugin. Packages not named main are ignored.

# ldflagis reference
#
# -s Omit the symbol table and debug information (makes binary smaller)
# -w Omit the DWARF symbol table (makes binary smaller
# -g Disable Go Package data checks (faster compilation)
# -d Disable Gneration of Dynamic Executables (makes a more unversal binary i believe?)

#******
echo ""
echo "            ***********************************************"
echo "            ***       Terry's ULTIMATE GO 'Piler        ***"
echo "            ***               ${VERSION}                     ***"
echo "            ***********************************************"
echo ""
tput sgr0
echo "                 (Make sure you are in the directory)"
echo "           (of the GO program or Module you want to compile)"
echo ""


CURR_ARCH=`uname -m`

GOARCH_BIN_ARCH=amd64              # Architecture type 64bit or 32bit  (amd64 or 386)

ALT_BIN_NAME=""     # used if -o was specifiedx
ALT_DEPLOY_DIR=""   # used if -d was specified

GOOS_BIN_TYPE=darwin             # The binary "type"  ..can be one of: win | mac | linux | win32 | linux32      
BIN_EXT='.mac'              # the exension.. this is .mac for mac binarues, .linux for linux and .exe for windows

PURGE_ALL_CACHES="no"       # We dont purge Go Caches by default....only if NECESSARY
CURR_DIR_NAME=${PWD##*/}        # Current directory name ... MINUS this path
BIN_NAME=${CURR_DIR_NAME}         # The name of the binary...  output filename.. this defaults to the current name of the directory you are in

RUN_MODE=""       # Can also be initmod or testmod
JUST_CLEAN=""
USE_GIT_TAG_for_VERSION=""      # determines the behavior of the version that is embedded in the binary.. defaults to just teh COMMIT
USE_FORCE_VERSION=""

# = = = = = = = = = = = = FUNCTIONS Start Here
GO_CACHE_PURGE() {

    ECHO_COLOR 6 "       **"
    ECHO_COLOR 6 "       ****  PURGING ALL GO and GO_MODULE CACHES "    
    ECHO_COLOR 6 "       **"


    ECHO_COLOR 3 "  WARNING: Purging is a MAJOR operation and will force you to reset everything (including VS CODE)"
    ECHO_COLOR 3 "  You might want to do:" nobr
    ECHO_COLOR 2 "  gobuild --clean"
    ECHO_COLOR 5 "  Are you SURE?: " nobr
    RESPONSE=""
    read -e -p "Type YES: " RESPONSE

    if [[ ! $RESPONSE == *"YES"* ]]; then
        echo ""
        exit -9
    fi


    sudo chmod 777 ~/go
    go clean --cache
    go clean --modcache

    # removes go.mod and sum from curr dir
    sudo rm -rf go.mod
    sudo rm -rf go.sum

    # prevents erran g.mod/sum in home
    sudo rm -rf ~/go.mod
    sudo rm -rf ~/go.sum

    sudo rm -rf ~/Library/Caches/go-build
    sudo rm -rf ~/Library/Application\ Support/go
    #sudo rm -rf ~/go

    ECHO_COLOR 3 "       **"
    ECHO_COLOR 3 "       ****  PURGE COMPLETE!!! "    
    ECHO_COLOR 3 "       **"    

}


DO_QUICK_COMMIT() {

    # Reference:
    #      https://stackoverflow.com/questions/26042390/git-add-asterisk-vs-git-add-period

    MESSAGE="Auto QuickCommit - "
    if [[ ! -z ${1} ]]; then
        MESSAGE=${1}
    fi

    echo ""
    ECHO_COLOR 2 " = = COMMITTING to GIT REPO = = "
    echo ""    

    mdate=`date +"%A - %b %d, %Y @ %H:%M"`
    COMMIT_MESS="${MESSAGE} | ${mdate}"

    # star really has no special meaning  | git add *   and only applies to cur dir
    git add -A    # gets everything (including deleted files
    git add .     # gets everything recursevly   (but skips deleted)

    git add --all
    git add -u
    
    git commit -m "${COMMIT_MESS}" .
    git push -f

}




BASE_REPO=""
BASE_MODULE_NAME=""
REPO_URL=""
FULL_GIT_MOD_NAME=""


SHOW_REPO_INFO() {
    TEMP_BASE=`basename $PWD`
    echo ""
    ECHO_COLOR 7 "  REPO: " nocr
    ECHO_COLOR 3 ${REPO_URL}
    ECHO_COLOR 7 "  Official GO MOD Name: " nocr   
    ECHO_COLOR 6 "${FULL_GIT_MOD_NAME}"
    echo ""
    ECHO_COLOR 2 "  import " nocr
    ECHO_COLOR 6 ". \"${FULL_GIT_MOD_NAME}\""
    echo "      or  "
    ECHO_COLOR 2 "  import " nocr
    ECHO_COLOR 6 "${TEMP_BASE} \"${FULL_GIT_MOD_NAME}\""    
    echo ""    
    
}

# This provides support for making and testing GO MODULES
# This is also does what the former quickCommit.sh used to do
GO_MODULES_SUPPORT_ENGINE() {
    echo ""
    ECHO_COLOR 3 "= = = = = = ="
    ECHO_COLOR 3 "= = = = = = =  GO MODULE Engine"
    ECHO_COLOR 3 "= = = = = = ="

    #all of this is used to get the proper go module name that we use later in go mod init
    #github.com/acedev0/LEGACY/Gadgets

    CHECK_for_REPO_NAME=`git remote show origin 2>&1 | grep Fetch`
    BASE_REPO=$(basename `git rev-parse --show-toplevel`)
    BASE_MODULE_NAME=$(echo $PWD | grep -Eo ${BASE_REPO}'(...).*')
    REPO_URL=`echo $CHECK_for_REPO_NAME | cut -d \@ -f 2`
    FULL_GIT_MOD_NAME=""

# = = = = = = = = =
# = = = = = = = = = Only spports github.com now... Will need to add support for Gitlab, AzureDevops, and Bitbucket later
# = = = = = = = = =
# = = = For GITHUB.com format GO Modules
IFS=/ read -a foo <<< "${REPO_URL}"
URL=${foo[0]}
GIT_USER=${foo[1]}
FULL_GIT_MOD_NAME="${URL}/${GIT_USER}/${BASE_MODULE_NAME}"

# echo "CheckRepoNAME: ${CHECK_for_REPO_NAME}"
# echo "REPO_URL: ${REPO_URL}"
# echo "Base REPO: ${BASE_REPO}"
# echo "BASE_MODULE_NAME: ${BASE_MODULE_NAME}"
# echo "FULL GITMOD:"
# ECHO_COLOR 2 $FULL_GIT_MOD_NAME


    if [[ $FULL_GIT_MOD_NAME == "" ]]; then
        echo ""
        ECHO_COLOR 5 " ERROR! Cant Generate proper GOMOD name for Git Module!!!"
        echo ""
        exit
    fi

    if [[ $CHECK_for_REPO_NAME == "" ]]; then
        echo ""
        echo ${CHECK_for_REPO_NAME}
        echo ""
        ECHO_COLOR 5 " ERROR! Cant find the .git repo FOLDER!!!"
        ECHO_COLOR 3 " You need to already be in a GIT repo for this to work properly!!!!"
        echo ""
        exit
    fi

    if [[ ${RUN_MODE} == *"repo"* || ${RUN_MODE} == *"modu"* || ${RUN_MODE} == *"gomod"* ]]; then

        SHOW_REPO_INFO
        exit
    fi
    

    echo ""
    ECHO_COLOR 3 "= = = Testing Module to make sure it compiles.."
    echo ""

    rm go.*
    go clean --modcache
    go clean --testcache
    go clean --cache
    go mod init ${FULL_GIT_MOD_NAME}
    go mod tidy

    testCompile=`go test 2>&1`
    if [[ $testCompile == *"no test files"* ]]; then
        echo ""
        ECHO_COLOR 2 " = = This MODULE COMPILED SUCCESSFULLY!!! = = "
        echo ""

    else 
        ECHO_COLOR 5 " ERROR in MOD Compile!"
        echo ""
        ECHO_COLOR 3 ${testCompile}
        echo ""

	    go test
	    echo ""
        exit
    fi



    #3. And finally... unless you are just testing the module, lets commit this to the GIT REPO automatically:
    if [[ ${RUN_MODE} == *"testmod"* ]]; then
	ECHO_COLOR 3 " Only TESTING Module...so exitting.."

        SHOW_REPO_INFO
    	exit
    fi
        
    DO_QUICK_COMMIT "MODULE update for GO"


    SHOW_REPO_INFO
}





DO_Light_CLEAN() {

    ECHO_COLOR 3 "= = = = = = ="
    ECHO_COLOR 3 "= = = = = = =  GO Light Clean" nobr
    ECHO_COLOR 2 " Done! "
    ECHO_COLOR 3 "= = = = = = ="
    echo ""
    echo ""

    rm -rf go.mod
    
}
# = = = = = = = = = = = = end of HELPER FUNCTIONS


# New for supporting A1 Macs 
if [[ ($CURR_ARCH == *'arm64'*) ]]
then
	GOARCH_BIN_ARCH="arm64"
fi


#1. First we iterate through the parameters that are passed to figure out how we need to "build" this binary
for param in "${@:1}"   #"$@"    
do
    #1b. First increment the parameter count (starts at 0)
    let "pnum=pnum+1"

    #1c. Now split out the parameter value
    p_temp=${@:pnum+1}
    nsplit=(${p_temp// / })
    p_value=${nsplit[0]}

    #1d. Get setup to build for linux
    if [[ ($param == linux) || ($param == lin) ]]; then
        echo "    --- Building for LINUX platform "
        BIN_EXT='.LINUX'
        GOOS_BIN_TYPE=linux
        GOARCH_BIN_ARCH="amd64" 

    elif [[ $param == *'linux64'* || $param == *'linuxARM' ]]; then
        echo "    --- Building for LINUX ARM 64bit platform "
        BIN_EXT='.LINUX_arm64'
        GOOS_BIN_TYPE=linux
        GOARCH_BIN_ARCH="arm64" 

    #1e. To get a CUSTOM name for the output file (we wont append an extension)
    elif [[ $param == *'-name'* ]]; then
        echo ""
        echo " ALTNAME IS: ${p_value}"
        echo ""
        ALT_BIN_NAME=${p_value}

    elif [[ ($param == *'apple'*)  ]]; then
        echo "    --- Building for APPPLE Silicone!!! (new) "
        GOARCH_BIN_ARCH="arm64"

    elif [[ ($param == *'intel'*) ]]; then
        echo "    --- Building for Mac INTEL binary "
	    GOARCH_BIN_ARCH="amd64"
	    BIN_EXT=".macINTEL"

    elif [[ ($param == *'-d'*) || ($param == *'-dest'* ) ]]; then

        ALT_DEPLOY_DIR=${p_value}        
    
    #1f. Get setup to build binary for windows
    elif [[ ($param == *'win'*) ]]; then
        echo "    --- Building for WINDOWS platform"
        BIN_EXT='.exe'
        GOOS_BIN_TYPE=windows
        GOARCH_BIN_ARCH=amd64

    #1g. Get Override, in case we need a 32bit binary
    elif [[ ($param == 32bit) || ($param == 32) ]]; then
        echo "    --- Generating 32bit Binary ---"
        GOARCH_BIN_ARCH=386

    # 1h. Thisis override behavior. If you want to force using the GIT TAG for versioning the binary
    #     (assuming there is one)...specify --tag || --tagver
    elif [[ ($param == *'tag'*) ]]; then
        USE_GIT_TAG_for_VERSION="yes"

    # 1i. This is to quickly commit the git repo we are in
    elif [[ ($param == *'-clean'*) ]]; then
        DO_Light_CLEAN

    # 1i. This is to quickly commit the git repo we are in
    elif [[ ($param == *'commit'*) ]]; then
        DO_QUICK_COMMIT

    elif [[ ($param == *'-forcever'*) ]]; then
        echo "    --- Baking with SPECIFIC version: ${p_value} "
        USE_FORCE_VERSION=${p_value}


# for doing go module stuff.. 
    elif [[ ($param == *'tmod'*) || ($param == *'savemod'* ) ]]; then
        RUN_MODE=${param}

    elif [[ ($param == *'repo'*) || ($param == *'mod'* ) ]]; then
        RUN_MODE=${param}        

    elif [[ ($param == *'-purge'*) ]]; then
        PURGE_ALL_CACHES="yes"

    # For getting HELP
    elif [[ ($param == *'-help'*) || -z ${1} ]]; then
cat <<EOMHELP
        Simple use!   (note, the commands can all be stacked) 
            gobuild             - (builds by default for whatever platform you are on)
            gobuild linux       - Compiles a binary for LINUX OS
            gobuild windows     - Compiles an EXE for windows
            gobuild macintel    - Builds explicitly for Mac Intel x86
            gobuild apple       - Builds explicitly for Apple Silicone (arm64 processor)

        Also these are helpful:

            gobuild --purge     - Will purge all your module and go cache. 
                                  Ensures you have latest version of module deps

            gobuild --clean     - Does a simply clean (in the current directory)
            gobuild --commit    - Commits to the GIT repo (if you are using one)
            gobuild -o [FNAME]  - Lets you specify the NAME of the output file you want go to generate
            gobuild -d [DIR]    - lets you specify the output PATH you want the go binary to be generated in
            gobuild --tagver    - If specified, will use a branch tag (if it exists) instead of the last COMMIT
                                  (NOTE: All binaries are stamped with a version, this just changes behavior)

            gobuild --forcever  - If you are a savage and want to use your OWN damn version, specify it here

        For Go MODULE Support
            gobuild initmod     - Will setup a proper GO Module and spit out the path 
                                  you need for importing (be sure you are in a GIT repo!!)

            gobuild testmod     - Will test the module to make sure it compiles.. useful before commiting

EOMHELP
        exit 0
    fi
done
# end of FOR loop to get the input params

# = = = = = = = = = = = = = = = = 
# = = = = ERROR HANDLING = = = = = 
# Error hanling.. make sure we have some .go files
if [ -n "$(ls -A *.go 2>/dev/null)" ]
then
  echo ""
else

    ECHO_COLOR 3 "          NO *.go files were found!!! "
    echo ""
    echo ""

  exit -9
fi


# = = = = = = = = = = = = = = = = 
# = = = = ERROR HANDLING = = = = = 

# If you are in a parent directory that is a symbolic link GO DOESNT LIKE that
# it wont be able to find your imports and packages properly
CHECK=`ls -ltd $PWD`
PARENT_DIR_SYMLINK=""
if [[ ${CHECK} == *"lrw"* || ${CHECK} == *"->"* ]]; then
    PARENT_DIR_SYMLINK="yes"

    echo ""
    echo ""
    ECHO_COLOR 5 "   ERROR! Symlink Parent Directory Found!!!"
    ECHO_COLOR 3 "   Looks like your current directory is actualy a SYMLINK!!"
    echo $CHECK
    echo ""
    ECHO_COLOR 3 "   Go doesnt work right if you build from a SymLink directory, and you'll see "
    ECHO_COLOR 3 "   errors about not being able to find your imports"
    ECHO_COLOR 3 "   Switch to a fully qualified path (not a symlink) and run your go build from there!!"
    echo ""
    exit -9
fi


# = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
# = = = = = = = = = = = = MAIN = = = = = = = = = = = = = =
# = = = = = = = = = = = = = = = = = = = = = = = = = = = = =



#8. Build binary for MAC by default!!..
if [[ ($GOOS_BIN_TYPE == darwin) ]]
then
    echo "         ** Building for MAC/OSX platform * *"
fi

#8b, just in case 32 bit was speicifed
if [[ ($GOARCH_BIN_ARCH == 386 )]]
then 
    BIN_NAME=$BIN_NAME"_32bit"
fi

#8c. For the output filename. we auto generate based on the current directory name
OUTPUT_FILE=$BIN_NAME$BIN_EXT
FULL_FILE=$OUTPUT_FILE

#9 Now.. if they specified a DIFFERENT NAME for the output file:
#You can also specify a FULL PATH
if [[ ! -z $ALT_BIN_NAME ]]; then
    OUTPUT_FILE=$ALT_BIN_NAME
    FULL_FILE=$OUTPUT_FILE

    justFILE="$(basename "${ALT_BIN_NAME}")"
    BIN_NAME=${justFILE}
    OUTPUT_FILE=$justFILE
fi

#9b. Also, if they specified a DIFFERENT dest directory for output
if [[ ! -z $ALT_DEPLOY_DIR ]]; then    
    sudo mkdir -p $ALT_DEPLOY_DIR
    sudo chmod 775 $ALT_DEPLOY_DIR
    FULL_FILE="${ALT_DEPLOY_DIR}/${OUTPUT_FILE}"
fi


#9c a little error handling.. They can specify either a PATH in ALT_BIN_NAME with --out
#   Or they can specify a dest path with --dest  (in ALT_DEPLOY_DIR)
#   BUT NOT BOTH!!! if both are detected with a path, we default to --out and ignore ALT_DEPLOY_DIR
if [[ ( ! -z $ALT_DEPLOY_DIR) && ($ALT_BIN_NAME == *'/'*) ]]; then
    ALT_DEPLOY_DIR=""
    justDIR="$(dirname "${ALT_BIN_NAME}")"
    justFILE="$(basename "${ALT_BIN_NAME}")"
    sudo mkdir -p $justDIR
    sudo chmod 775 $justDIR    

    OUTPUT_FILE=$justFILE
    BIN_NAME=${justFILE}

    FULL_FILE=${ALT_BIN_NAME}
fi

#echo "BIN_NAME is: ${BIN_NAME}"
# echo "OUTPUT_FILE: ${OUTPUT_FILE}"
# echo "ALT_DEPLOY_DIR: ${ALT_DEPLOY_DIR}"
# echo "FULL_FILE: ${FULL_FILE}"
# exit 


#12. -=-=-= MAIN Action starts here!! -=-=-=-=-

#12b. Now lets check if we are purging any CACHES
if [[ ${PURGE_ALL_CACHES} == "yes" ]]; then
    GO_CACHE_PURGE
fi

# Light clean just removes a go.mod which go.tidy re-pulls
if [[ ${LIGHT_CLEAN} == "yes" ]]; then
    DO_Light_CLEAN
fi

sudo chmod 777 ~/go
# Default Run mode (is to compile GO Programs)

# If runmode is initmod or testmod... testmod just compiles to test... initmod also commits to git
if [[ ${RUN_MODE} == *"tmod"* ]]; then

    GO_MODULES_SUPPORT_ENGINE


elif [[ ${RUN_MODE} == *"repo"* || ${RUN_MODE} == *"mod"* ]]; then

    GO_MODULES_SUPPORT_ENGINE

elif [[ -z ${RUN_MODE} ]]; then

    # First do some cleanup
    sudo rm -rf $FULL_FILE
    sudo rm -rf *.LINUX
    sudo rm -rf *.mac
    sudo rm -rf *.exe    

    # We need to always run go mod and gomod tidy.. which downloads all module dependencies for this GO Program    
    # cheap way to run and ignore errors

    # using NEW organization method.. all "programs" will be called local

    go mod init "local"
    go mod tidy

    # = = = = = = = == = = == =
    # = = ==   Stamped Versioning in Binary
    # = = = = = = = == = = == =
    # This is just in case we arent in a GIT repo.. or no custom version was specified
    HAVE_GIT=`which git 2>&1`
    
    VERSION_to_USE=""
    gdate=`date +"%m%d%Y.%H.%M%S"`
    GHETTO_DEFAULT_VER="1.${gdate}"
    
    #2b. if some crazy PPL want to use their own custom version, instead of the uber cool auto version logic
    if [[ ! -z ${USE_FORCE_VERSION} ]]; then
        VERSION_to_USE="${USE_FORCE_VERSION}"

    elif [[ -z ${HAVE_GIT} || ${HAVE_GIT} == *'fatal'* || ${HAVE_GIT} == *'command not found'* ]]; then

        VERSION_to_USE=${GHETTO_DEFAULT_VER}
    else
        
        #3 First we'll check to use the last git commit
        VERSION_to_USE=`git rev-parse --short=6 HEAD 2>&1`

        if [[ -z ${VERSION_to_USE} || ${VERSION_to_USE} == *'fatal'* || ${VERSION_to_USE} == *'command not found'* ]]; then
            VERSION_to_USE=${GHETTO_DEFAULT_VER}
        else
            #3b, Howeveer if USE_GIT_TAG_for_VERSIOn was specified, we use the TAG (if there is one)
            if [[ ! -z ${USE_GIT_TAG_for_VERSION} ]]; then    
                TAGVER=`git tag -l --sort=refname`

                # If we HAVE a tag, replace VERSION_to_USE with it
                if [[ ! -z ${TAGVER} ]]; then
                    VERSION_to_USE=$TAGVER
                fi        
            fi
        fi
    fi

    

    #5. Next we show the command we will use to compile go with
    # ref: https://icinga.com/blog/2022/05/25/embedding-git-commit-information-in-go-binaries/

    #TEXT_OF_FULL_EXEC_COMMAND="GOOS=$GOOS_BIN_TYPE GOARCH=$GOARCH_BIN_ARCH go build -ldflags=\"-s -w -X main.VERSION_NUM=$($TPUTBIN setaf 2)${VERSION_to_USE}$($TPUTBIN setaf 3)\" -buildmode=exe -o $($TPUTBIN setaf 2)${FULL_FILE} $($TPUTBIN setaf 3)./..."
    #GOOS=$GOOS_BIN_TYPE GOARCH=$GOARCH_BIN_ARCH go build -ldflags="-s -w -X main.VERSION_NUM=${VERSION_to_USE}" -buildmode=exe -o $FULL_FILE ./... 2>&1

    TEXT_of_FULL_COMMAND="GOOS=$GOOS_BIN_TYPE GOARCH=$GOARCH_BIN_ARCH go build -ldflags=\"-s -w -X main.VERSION_NUM=${VERSION_to_USE}\" -buildmode=exe -o ${FULL_FILE}" 
    #GOOS=$GOOS_BIN_TYPE GOARCH=$GOARCH_BIN_ARCH go build -ldflags="-s -w -X main.VERSION_NUM=${VERSION_to_USE}" -buildmode=exe -o $FULL_FILE 

    echo ""
    ECHO_COLOR 6 "          NOTE: Here is the go build i will use: "
    echo ""
    ECHO_COLOR 3 " ${TEXT_of_FULL_COMMAND} " 
    echo ""

    eval ${TEXT_of_FULL_COMMAND} 2>&1 | tee /tmp/gobuild.log


    tput setaf 3    

    


    # easy error handling.. check if file exists
    echo " "
    if [[ -e ${FULL_FILE} ]]; then
        ECHO_COLOR 6 "          NOTE: The binary has been stamped with the version: " nobr
        ECHO_COLOR 2 "${VERSION_to_USE}"
        ECHO_COLOR 6 "                If you want to display this from within your GO program"
        ECHO_COLOR 6 "                just use " nobr
        ECHO_COLOR 7 "fmt.Println(" nobr
        ECHO_COLOR 2 VERSION_NUM nobr
        ECHO_COLOR 7 ")"
        echo ""

        tput setaf 2                           
        echo "           ***************************"
        echo "           **** Compile SUCCESS!! ****"
        echo "           ***************************"        
        echo ""
    else 

        ECHO_COLOR 5 "**** ERROR: Compile FAILURE ****"
        echo ""

        COMPILE_RESULT=`cat /tmp/gobuild.log`

        # If we have go get errors, usually means a module import is wrong
        if [[ ${COMPILE_RESULT} == *"no required module provides"* ]]; then

            echo ""
            ECHO_COLOR 3 "  WARNING: Looks like there are issues with some GO imports!"
            ECHO_COLOR 6 "  ( I will attempt to fix the issue, using 'go get' on the bad imports )"


            echo ""
            RESULTX="$(cat /tmp/gobuild.log)"
            RESULT="${RESULTX%x}"

            while IFS= read -r line
            do
                if [[ $line == *"go get"* ]]; then
                    GOGET="$(${line})"
                fi
            done <<< "$RESULT"

            echo ""
            ECHO_COLOR 3 "  If you are still getting errors, This usually means there is an import with the WRONG path"
            ECHO_COLOR 3 "  It could also mean the module is broken (or not created properly) and wont compile"
            echo ""
            ECHO_COLOR 3 "  NOTE: Worst case scenario, you can do a massive cache cleanup with: " 
            echo ""
            ECHO_COLOR 6 "     gobuild --clean"
            ECHO_COLOR 5 "        or"
            ECHO_COLOR 6 "     gobuild --purge"

     
            echo ""
            echo ""
            exit -9
        fi
    fi


    # Also, if this was a windows compile, lets ZIP the file so we can distribute it to Windows People
    if [[ ${GOOS_BIN_TYPE} == "windows" ]]; then
        echo ""
        ZIP_FILE_NAME="${BIN_NAME}.zip"     # this is only for WIN EXEs
        FULL_ZIP="${DEPLOY_DIR}/${ZIP_FILE_NAME}"

        ECHO_COLOR $white "Now creating ZIP file (for Windows People): " nobr
        ECHO_COLOR $green ${FULL_ZIP}        
        zip ${FULL_ZIP} ${FULL_FILE}

        # and remove the original unzipped full file.. avoid confusion
        echo ""        
        rm -rf ${FULL_FILE}
    fi

fi

echo ""
tput sgr0

